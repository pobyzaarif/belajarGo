package protobuf

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func LearnProto() {
	socMed := &Profile_SocialMedia{
		Twitter:   "pobyzaarif",
		Instagram: "pobyzaarif",
	}

	profile := &Profile{
		Name:        "poby",
		Age:         1,
		Married:     true,
		Salary:      1.234,
		SocialMedia: socMed,
		Hobbies:     []string{"coding", "singing"},
	}

	// fmt.Println(profile)

	data, err := proto.Marshal(profile)
	if err != nil {
		log.Fatal("err marshal : ", err)
	}

	fmt.Println(data)

	newProfile := &Profile{}
	err = proto.Unmarshal(data, newProfile)
	if err != nil {
		log.Fatal("err unmarshal : ", err)
	}

	fmt.Println(newProfile.GetSocialMedia())
	fmt.Println(newProfile.SocialMedia)
}
