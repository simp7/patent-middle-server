package main

type NLP interface {
	Process(string) ([]byte, error)
}
