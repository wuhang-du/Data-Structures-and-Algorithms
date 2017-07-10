//
// date: 2017-7-10
// desc: hello
//

func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    if haystack == "" {
        return -1
    }
    lh := len(haystack)
    ln := len(needle)
    check := needle[0]
    start := 0
    
    for start + ln <= lh {
        if string(haystack[start:start+ln]) == needle {
            return start
        }
        start++
        for start < lh && haystack[start] != check {
            start++
        }
    }
    return -1
}
