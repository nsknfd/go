package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	u1 := uuid.NewV1()
	fmt.Printf("UUIDv1: %s\n", u1)
	u2 := uuid.NewV2(uuid.DomainPerson)
	fmt.Printf("UUIDv2: %s\n", u2)
	u3 := uuid.NewV3(uuid.NamespaceURL, "https://github.com/nsknfd/go")
	fmt.Printf("UUIDv3: %s\n", u3)
	u4 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u4)
	u5 := uuid.NewV5(uuid.NamespaceURL, "https://github.com/nsknfd/go")
	fmt.Printf("UUIDv5: %s\n", u5)
}
