func checkInclusion(s1 string, s2 string) bool {
    n, m := len(s1), len(s2)
    if n > m {
        return false
    }

    var arrS1, arrS2 [26]int
    for i := 0; i < n; i++ {
        arrS1[s1[i]-'a']++
        arrS2[s2[i]-'a']++
    }

    matches := 0
    for i := 0; i < 26; i++ {
        if arrS1[i] == arrS2[i] {
            matches++
        }
    }

    l := 0
    for r := n; r < m; r++ {
        if matches == 26 {
            return true
        }

        rIdx := s2[r] - 'a'
        if arrS1[rIdx] == arrS2[rIdx] {
            matches-- 
        }
        arrS2[rIdx]++
        if arr1 := arrS1[rIdx]; arr1 == arrS2[rIdx] {
            matches++ 
        }

        lIdx := s2[l] - 'a'
        if arrS1[lIdx] == arrS2[lIdx] {
            matches-- 
        }
        arrS2[lIdx]--
        if arr1 := arrS1[lIdx]; arr1 == arrS2[lIdx] {
            matches++ 
        }

        l++
    }

    return matches == 26
}