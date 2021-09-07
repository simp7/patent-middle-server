package files

import (
	"fmt"
	"github.com/simp7/patent-middle-server/config"
	"github.com/simp7/patent-middle-server/model"
	"gopkg.in/yaml.v3"
	"io"
	"io/fs"
	"os"
	"os/exec"
)

type system struct {
	ReadWrite
	skelFS ReadOnly
}

func System(main ReadWrite, sub ReadOnly) (sys *system, err error) {

	sys = &system{ReadWrite: main, skelFS: sub}

	if sys.isFirstTime() {

		fmt.Println("It is first time to run server.")
		fmt.Println("It will take few minutes, so BE PATIENT.")

		if err = sys.installEssentials(); err != nil {
			return
		}

	}

	if !sys.isLatest() {
		err = sys.update()
	}

	return

}

func (s *system) OpenLogfile() (*os.File, error) {
	return s.Open(LOG)
}

func (s *system) SaveCSVFile(group *model.CSVGroup) (file *os.File, err error) {

	filePath := New(group.ID + ".csv")

	if file, err = s.Open(filePath); err != nil {
		return
	}
	defer file.Close()

	if _, err = file.WriteString("name" + group.Separator + "item" + "\n"); err != nil {
		return
	}

	for _, v := range group.Data {
		_, _ = file.WriteString(v.Serialize(group.Separator))
	}

	return

}

func (s *system) RemoveCSVFile(group *model.CSVGroup) error {
	filePath := New(group.ID + ".csv")
	return s.Remove(filePath)
}

func (s *system) LoadConfig() (conf config.Config, err error) {

	file, err := s.Open(CONFIG)
	if err != nil {
		return config.Config{}, err
	}

	return s.decodeConfig(file)

}

func (s *system) Word2vec(args ...string) *exec.Cmd {
	return s.Command(WORD2VEC, args...)
}

func (s *system) LDA(args ...string) *exec.Cmd {
	return s.Command(LDA, args...)
}

func (s *system) LSA(args ...string) *exec.Cmd {
	return s.Command(LSA, args...)
}

func (s *system) isLatest() bool {

	realConf, _ := s.LoadConfig()
	skelConf, _ := s.skelConfig()

	return realConf.Version == skelConf.Version

}

func (s *system) updateVersion() error {

	realConf, _ := s.LoadConfig()
	skelConf, _ := s.skelConfig()

	realConf.Version = skelConf.Version
	return s.saveConfig(realConf)

}

func (s *system) skelConfig() (config.Config, error) {

	file, err := s.skelFS.Open(CONFIG)
	if err != nil {
		return config.Config{}, err
	}

	return s.decodeConfig(file)

}

func (s *system) decodeConfig(file io.ReadCloser) (config config.Config, err error) {

	defer file.Close()
	err = yaml.NewDecoder(file).Decode(&config)

	return

}

func (s *system) saveConfig(conf config.Config) (err error) {

	var data []byte

	if data, err = yaml.Marshal(conf); err == nil {
		err = s.Write(CONFIG, data)
	}

	return

}

func (s *system) initialize() (err error) {

	if s.isFirstTime() {

		fmt.Println("It is first time to run server.")
		fmt.Println("It will take few minutes, so BE PATIENT.")

		if err = s.installEssentials(); err != nil {
			return
		}

	}

	if !s.isLatest() {
		err = s.update()
	}

	return

}

func (s *system) update() (err error) {
	if err = s.updateVersion(); err == nil {
		err = s.updateFiles()
	}
	return
}

func (s *system) installEssentials() (err error) {

	if err = s.initFiles(); err != nil {
		return
	}

	fmt.Println("You should put password for sudo command to install/upgrade essential environment.")
	err = s.Command(INITIALIZE).Run()
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
	var created *os.File

	if s.IsExist(file) {
		return
	}

	if skelFile, err = s.skelFS.Open(file); err != nil {
		return
	}
	defer skelFile.Close()

	if created, err = s.Create(file); err != nil {
		return
	}
	defer created.Close()

	if _, err = io.Copy(created, skelFile); err != nil && err != io.EOF {
		return
	}

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
