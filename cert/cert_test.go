package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Box", "2025-08-28")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert data should not be nil.")
	}
	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. Expected='GOLANG COURSE, got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Box", "2025-28-08")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "aslkdjasdlkajsd lkasjdalksdjaklsdjasdkljasdklajsdalksdj"
	_, err := New(course, "Box", "2025-28-08")
	if err == nil {
		t.Errorf("Error should be returned on an too long course (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2025-28-08")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "aslkdjasdlkajsdlkasjdalksdjaklsdjasdkljasdklajsdalksdjsadasdlkasjdalskdjaslkdjaskldjaslkdjaskldj"
	_, err := New("Golang", name, "2025-28-08")
	if err == nil {
		t.Errorf("Error should be returned on an too long course (course=%s)", name)
	}
}

func TestDateValid(t *testing.T) {
	_, err := New("Golang", "Box", "2025-08-28")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
}

func TestDateInvalid(t *testing.T) {
	invalid := []string{
		"",           // empty
		"not-a-date", // garbage
		"2025/08/28", // wrong separators
		"2025-13-01", // bad month
		"2025-00-10", // bad month
		"2025-02-30", // bad day
		"2023-02-29", // non-leap year
		"28-08-2025", // wrong layout
	}

	for _, in := range invalid {
		if _, err := parseDate(in); err == nil {
			t.Errorf("parseDate(%q) expected error, got nil", in)
		}
	}
}
