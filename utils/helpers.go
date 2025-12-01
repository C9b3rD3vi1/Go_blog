package utils

import (
	"strings"
    "encoding/json"
    "github.com/C9b3rD3vi1/Go_blog/models"
    
)

// Template helper functions
func HasTechStack(project *models.Projects, techStackID uint) bool {
    for _, ts := range project.TechStacks {
        if ts.ID == techStackID {
            return true
        }
    }
    return false
}

// Template function to parse JSON (for gallery display)
func ParseJSON(s string) []string {
    var result []string
    if err := json.Unmarshal([]byte(s), &result); err != nil {
        return []string{}
    }
    return result
}

func SplitString(s string, sep string) []string {
    return strings.Split(s, sep)
}

func Add(a int, b int) int {
    return a + b
}


func Trim(s string) string {
    return strings.Trim(s, " ")
}
