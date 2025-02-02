package handler

import "log"

func HandleError(err error, doneChan chan bool) {
	if err != nil {
		log.Fatal(err);
		if doneChan != nil {
			doneChan <- true;
		}
		return
	}
}