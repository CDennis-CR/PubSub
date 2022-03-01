package pubsub

import (
	"fmt"
	"log"
	"sync"
)

//getReady creates empty pubsub and users instances to begin the application
//
//Boots in the mux
func getReady(superUsername, superUserpassword string) *PubSub {
	//generate special user `ping`
	superUserPing, err := createNewUser(superUsername, superUserpassword)
	if err != nil {
		log.Fatalln(err)
	}
	//echo superuser login to std.out
	log.Printf("Superuser Ping created.\nUUID: %s", superUserPing.UUID)
	//new core
	users := Users{superUserPing.UsernameHash: superUserPing}
	pubsub := &PubSub{
		Topics: make(Topics),
		Users:  users,
		mu:     &sync.RWMutex{},
	}
	//start regular task ticks
	go pubsub.metranome()
	//start persistance layer
	pubsub.persistLayer, err = NewUnderwriter(pubsub)
	if err != nil {
		log.Fatalln(fmt.Errorf("Error spinning up new Underwriter object: %v", err))
	}
	pubsub.persistLayer.Launch()

	return pubsub
}
