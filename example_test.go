// SPDX-FileCopyrightText: 2012 Jeramey Crawford <jeramey@antihe.ro>
// SPDX-License-Identifier: BSD-2-Clause

package crypt_test

import (
	"fmt"

	"github.com/tredoe/crypt"
	_ "github.com/tredoe/crypt/sha256_crypt"
)

func ExampleCrypt() {
	crypt := crypt.SHA256.New()
	ret, _ := crypt.Generate([]byte("secret"), []byte("$5$salt"))
	fmt.Println(ret)

	err := crypt.Verify(ret, []byte("secret"))
	fmt.Println(err)

	// Output:
	// $5$salt$kpa26zwgX83BPSR8d7w93OIXbFt/d3UOTZaAu5vsTM6
	// <nil>
}
