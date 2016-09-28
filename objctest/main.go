package main

/*
#include <Foundation/Foundation.h>
#include <Foundation/NSObjCRuntime.h>
#include <IOKit/IOKitLib.h>

#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework IOKit

void hello() {
    NSLog(@"Hello World from Objective-C!");
}

const char* NSStringToChar(NSString *s) {
    if (s == NULL)
        return NULL;

    return [s UTF8String];
}

const char* getUUID() {
    CFTypeRef cfSerialNum;
    io_service_t platExpert;

    platExpert = IOServiceGetMatchingService(kIOMasterPortDefault, IOServiceMatching("IOPlatformExpertDevice"));
    if (!platExpert)
        return "";

    cfSerialNum = IORegistryEntryCreateCFProperty(platExpert, CFSTR(kIOPlatformUUIDKey), kCFAllocatorDefault, 0);
    if (!cfSerialNum)
        return "";

    IOObjectRelease(platExpert);

    return NSStringToChar(cfSerialNum);
}
*/
import "C"

import "fmt"

func main() {
	C.hello()
	val := C.getUUID()
	fmt.Printf("Serial #: %s\n", C.GoString(val))
}
