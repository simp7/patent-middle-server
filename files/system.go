package files

import (
	"fmt"
	"github.com/simp7/patent-middle-server/config"
	"github.com/simp7/patent-middle-server/model"
	"gopkg.in/yaml.v3"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"sync"
)

var once sync.Once

type system struct {
	ReadWrite
	skelFS ReadOnly
	config config.Config
}

func System(main ReadWrite, sub ReadOnly) (sys *system, err error) {

	sys = &system{ReadWrite: main, skelFS: sub}

	if err = sys.initialize(); err != nil {
		log.Fatal(err)
	}

	if _, err = sys.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	err = sys.update()

	return

}

func (s *system) BindLogFiles(logPath ...string) (io.WriteCloser, error) {

	result := make([]io.WriteCloser, 0)

	if file, err := s.Open(LOG, true); err == nil {
		result = append(result, file)
	}

	for _, logFile := range logPath {

		file, err := s.Open(New(logFile), true)
		if err != nil {
			return nil, err
		}

		result = append(result, file)

	}

	return s.bindLogger(result...), nil

}

func (s *system) SaveCSVFile(group *model.CSVGroup) (err error) {

	filePath := New(group.FileName)
	stream := make(chan string)

	go func() {
		if err = s.WriteWithChan(filePath, stream); err != nil {
			return
		}
	}()

	if err != nil {
		return
	}

	stream <- group.Header()
	wg := sync.WaitGroup{}

	for _, data := range group.Data {
		wg.Add(1)
		go func(line string) {
			stream <- line
			wg.Done()
		}(data.Serialize(group.Separator))
	}
	wg.Wait()

	close(stream)
	return

}

func (s *system) RemoveCSVFile(group *model.CSVGroup) error {
	filePath := New(group.FileName)
	return s.Remove(filePath)
}

func (s *system) LoadConfig() (conf config.Config, err error) {

	once.Do(func() {

		var file *os.File

		if file, err = s.Open(CONFIG, false); err != nil {
			return
		}
		if s.config, err = s.decodeConfig(file); err != nil {
			return
		}

		err = file.Close()

	})

	conf = s.config
	return

}

func (s *system) Word2vec(args ...string) ([]byte, error) {
	return s.Command(WORD2VEC, args...).CombinedOutput()
}

func (s *system) LDA(args ...string) ([]byte, error) {
	return s.Command(LDA, args...).CombinedOutput()
}

func (s *system) LSA(args ...string) ([]byte, error) {
	return s.Command(LSA, args...).CombinedOutput()
}

func (s *system) update() (err error) {

	skelConf, _ := s.skelConfig()
	if s.config.Version == skelConf.Version {
		return
	}

	if err = s.updateVersion(skelConf.Version); err == nil {
		err = s.updateFiles()
	}

	return

}

func (s *system) updateVersion(target string) error {
	s.config.Version = target
	return s.saveConfig()
}

func (s *system) skelConfig() (config.Config, error) {

	file, err := s.skelFS.Open(CONFIG)
	if err != nil {
		return config.Config{}, err
	}

	return s.decodeConfig(file)

}

func (s *system) decodeConfig(file io.ReadCloser) (config config.Config, err error) {
	err = yaml.NewDecoder(file).Decode(&config)
	return
}

func (s *system) saveConfig() (err error) {

	var data []byte

	if data, err = yaml.Marshal(s.config); err == nil {
		err = s.Write(CONFIG, data)
	}

	return

}

func (s *system) initialize() (err error) {

	fmt.Println("It is first time to run server.")
	fmt.Println("It will take few minutes, so BE PATIENT.")

	if err = s.installEssentials(); err != nil {
		return
	}
	err = s.update()

	return

}

func (s *system) installEssentials() (err error) {

	if err = s.initFiles(); err != nil {
		return
	}

	fmt.Println("You should put password for sudo command to install/upgrade essential environment.")
	if err = s.Command(INITIALIZE).Run(); err != nil {
		fmt.Println("error while executing " + INITIALIZE)
		return
	}
	fmt.Println("Installing/Upgrading process has been done! Good luck!")
	return

}

func (s *system) isFirstTime() bool {
	return !s.IsExist(ROOT)
}

func (s *system) initFiles() (err error) {

	_ = s.Mkdir(ROOT)

	files, err := s.skelFS.ReadDir(ROOT)
	if err != nil {
		return
	}

	for _, file := range files {
		fileName := New(file.Name())
		if err = s.copy(fileName); err != nil {
			return
		}
	}

	return

}

func (s *system) copy(file Path) (err error) {

	var skelFile fs.File

	if s.IsExist(file) {
		return
	}

	if skelFile, err = s.skelFS.Open(file); err != nil {
		return
	}
	defer skelFile.Close()

	err = s.Copy(file, skelFile)

	return

}

func (s *system) updateFiles() (err error) {

	var list []string

	if list, err = s.GetUpdateList(); err != nil {
		return
	}

	for _, file := range list {
		if err = exec.Command("rm", "-rf", file).Run(); err != nil {
			return
		}
	}

	err = s.installEssentials()
	return

}

func (s *system) bindLogger(logFiles ...io.WriteCloser) io.WriteCloser {

	var files logWriteCloser = make([]io.WriteCloser, 0)
	files = append(files, logFiles...)

	return files

}
