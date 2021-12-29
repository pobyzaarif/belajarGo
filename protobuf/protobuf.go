package protobuf

import (
	"fmt"
	"log"

	localProto "github.com/pobyzaarif/belajarGo/protobuf/proto"

	"google.golang.org/protobuf/proto"
)

func LearnProto() {
	socMed := &localProto.Profile_SocialMedia{
		Twitter:   "pobyzaarif",
		Instagram: "pobyzaarif",
	}

	profile := &localProto.Profile{
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

	newProfile := &localProto.Profile{}
	err = proto.Unmarshal(data, newProfile)
	if err != nil {
		log.Fatal("err unmarshal : ", err)
	}

	fmt.Println(newProfile.GetSocialMedia())
	fmt.Println(newProfile.SocialMedia)
}
