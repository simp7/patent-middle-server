package main

type NLP interface {
	Process(file string, args ...string) ([]byte, error)
}
