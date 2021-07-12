package main

type NLP interface {
	Process(string) (string, error)
}
