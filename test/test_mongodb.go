package main

import test_mongodb "go-test/module/test-mongodb"

func main() {
	testMongodb := &test_mongodb.TestMongodb{}
	testMongodb.Handle()
}
