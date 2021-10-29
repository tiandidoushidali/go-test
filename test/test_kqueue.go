package main

import test_kqueue "go-test/module/test-kqueue"

func main() {
	testKqueue := &test_kqueue.TestKqueue{}
	testKqueue.Handle()
}
