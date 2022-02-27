package webauthn

import (
	"bytes"
	"github.com/fxamacker/webauthn"
	"testing"
)

const attResponse1 = `{
    "rawId": "Aad50Szy7ZFb8f7wdfMmFO2dUdQB8StMrYBbhJprTCJIKVdbIiMs9dAATKOvUpoKfmyh662ZsO1J5PQUsi9yKNumDR-ZD4wevDYZnwprytGf5rn6ydyxQQtBYPSwS8u23FdVBxBqHa8",
    "id": "Aad50Szy7ZFb8f7wdfMmFO2dUdQB8StMrYBbhJprTCJIKVdbIiMs9dAATKOvUpoKfmyh662ZsO1J5PQUsi9yKNumDR-ZD4wevDYZnwprytGf5rn6ydyxQQtBYPSwS8u23FdVBxBqHa8",
    "response": {
        "attestationObject": "o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YVjszHUM-fXe8fPTc7IQdAU8xhonRmZeDznRqJqecdVRcUNFYfOzo63OAAI1vMYKZIsLJfHwVQMAaAGnedEs8u2RW_H-8HXzJhTtnVHUAfErTK2AW4Saa0wiSClXWyIjLPXQAEyjr1KaCn5soeutmbDtSeT0FLIvcijbpg0fmQ-MHrw2GZ8Ka8rRn-a5-sncsUELQWD0sEvLttxXVQcQah2vpQECAyYgASFYIMG7Y3fOeGecLpfn7XF_sV4OTc41tsbEPSECGfCiK480IlggH9-qVehm6Gj25SyZau17mB5c0YoTWBZ8ngdEka4EqOY",
        "clientDataJSON": "eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoib0dvd2lrQVZHcnZ4Y01uck50ODlCY0dsWnIwVVUwVWxfSm82U0R5RXJrTSIsIm9yaWdpbiI6Imh0dHBzOi8vd2ViYXV0aG53b3Jrcy5naXRodWIuaW8iLCJjcm9zc09yaWdpbiI6ZmFsc2V9"
    },
    "getClientExtensionResults": {},
    "type": "public-key"
}`

func TestAttResponse1(t *testing.T) {

	att, err := webauthn.ParseAttestation(bytes.NewReader([]byte(attResponse1)))
	if err != nil {
		t.Error(err)
	}
	t.Logf("parsed att: %+v", att)

}
