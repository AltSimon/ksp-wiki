package wikitext_parser

import (
	"reflect"
	"testing"
)

func TestSymmetricalHeading(t *testing.T) {
	heading := "=== Heading 3 ==="
	value := isCorrectHeadingSuffixSize(heading, 3)
	wants := true
	if value != wants {
		t.Errorf("isCorrectHeadingSuffixSize(%q, 3) = %v, but should be %v", heading, value, wants)
	}
}

func TestSymmetricalHeading2(t *testing.T) {
	heading := "==== Wrong Heading 3 ==="
	value := isCorrectHeadingSuffixSize(heading, 3)
	wants := false
	if value != wants {
		t.Errorf("isCorrectHeadingSuffixSize(%q, 3) = %v, but should be %v", heading, value, wants)
	}
}

func TestSymmetricalHeading3(t *testing.T) {
	heading := "=== Wrong Heading 3 ===="
	value := isCorrectHeadingSuffixSize(heading, 3)
	wants := false
	if value != wants {
		t.Errorf("isCorrectHeadingSuffixSize(%q, 3) = %v, but should be %v", heading, value, wants)
	}
}

func TestHeadingNumber1(t *testing.T) {
	heading := "=== Normal 3 Heading ==="
	value := getHeadingNumber(heading)
	wants := 3
	if value != wants {
		t.Errorf("getHeadingNumber(%q) = %v, wants %v", heading, value, wants)
	}
}

func TestHeadingNumber2(t *testing.T) {
	heading := "Not a heading"
	value := getHeadingNumber(heading)
	wants := -1
	if value != wants {
		t.Errorf("getHeadingNumber(%q) = %v, wants %v", heading, value, wants)
	}
}

func TestHeadingNumber3(t *testing.T) {
	heading := "=== Assymetrical heading ===="
	value := getHeadingNumber(heading)
	wants := -1
	if value != wants {
		t.Errorf("getHeadingNumber(%q) = %v, wants %v", heading, value, wants)
	}
}

func TestHeadingNumber4(t *testing.T) {
	heading := "==== Assymetrical heading ==="
	value := getHeadingNumber(heading)
	wants := -1
	if value != wants {
		t.Errorf("getHeadingNumber(%q) = %v, wants %v", heading, value, wants)
	}
}

func TestHeadingNumber5(t *testing.T) {
	heading := "======== Heading too large ========"
	value := getHeadingNumber(heading)
	wants := 6
	if value != wants {
		t.Errorf("getHeadingNumber(%q) = %v, wants %v", heading, value, wants)
	}
}

func TestHeadingToMarkdown(t *testing.T) {
	heading := "=== Heading 3 ==="
	value := HeadingToMarkdown(heading)
	wants := "### Heading 3"
	if value != wants {
		t.Errorf("HeadingToMarkdown(%q) = %v, wants %v", heading, value, wants)
	}
}

func TestBigHeadingToMarkdown(t *testing.T) {
	heading := "== Heading 2 =="
	value := HeadingToMarkdown(heading)
	wants := "## Heading 2\n---"
	if value != wants {
		t.Errorf("HeadingToMarkdown(%q) = %v, wants %v", heading, value, wants)
	}
}

func TestParseHeading(t *testing.T) {
	lines := []string{"==A heading 2 =="}
	value := Parse(lines)
	wants := []string{"## A heading 2\n---"}
	if !reflect.DeepEqual(wants, value) {
		t.Errorf("Parse([]string{\"==A heading 2 ==\"}) = %v, but wants %v", value, wants)
	}
}
