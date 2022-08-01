// Code generated by goa v3.7.6, DO NOT EDIT.
//
// sample HTTP client CLI support package
//
// Command:
// $ goa gen github.com/is-hoku/goa-sample/webapi/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	studentc "github.com/is-hoku/goa-sample/webapi/gen/http/student/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `student (get-student|get-students|create-student)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` student get-student --student-number 3740542320` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		studentFlags = flag.NewFlagSet("student", flag.ContinueOnError)

		studentGetStudentFlags             = flag.NewFlagSet("get-student", flag.ExitOnError)
		studentGetStudentStudentNumberFlag = studentGetStudentFlags.String("student-number", "REQUIRED", "Student's unique number")

		studentGetStudentsFlags = flag.NewFlagSet("get-students", flag.ExitOnError)

		studentCreateStudentFlags    = flag.NewFlagSet("create-student", flag.ExitOnError)
		studentCreateStudentBodyFlag = studentCreateStudentFlags.String("body", "REQUIRED", "")
	)
	studentFlags.Usage = studentUsage
	studentGetStudentFlags.Usage = studentGetStudentUsage
	studentGetStudentsFlags.Usage = studentGetStudentsUsage
	studentCreateStudentFlags.Usage = studentCreateStudentUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "student":
			svcf = studentFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "student":
			switch epn {
			case "get-student":
				epf = studentGetStudentFlags

			case "get-students":
				epf = studentGetStudentsFlags

			case "create-student":
				epf = studentCreateStudentFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "student":
			c := studentc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-student":
				endpoint = c.GetStudent()
				data, err = studentc.BuildGetStudentPayload(*studentGetStudentStudentNumberFlag)
			case "get-students":
				endpoint = c.GetStudents()
				data = nil
			case "create-student":
				endpoint = c.CreateStudent()
				data, err = studentc.BuildCreateStudentPayload(*studentCreateStudentBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// studentUsage displays the usage of the student command and its subcommands.
func studentUsage() {
	fmt.Fprintf(os.Stderr, `Service is the student service interface.
Usage:
    %[1]s [globalflags] student COMMAND [flags]

COMMAND:
    get-student: 学籍番号から学生を取得する。
    get-students: 学籍番号で昇順にソートされた全ての学生を取得する。
    create-student: 学生を登録する。

Additional help:
    %[1]s student COMMAND --help
`, os.Args[0])
}
func studentGetStudentUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] student get-student -student-number UINT32

学籍番号から学生を取得する。
    -student-number UINT32: Student's unique number

Example:
    %[1]s student get-student --student-number 3740542320
`, os.Args[0])
}

func studentGetStudentsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] student get-students

学籍番号で昇順にソートされた全ての学生を取得する。

Example:
    %[1]s student get-students
`, os.Args[0])
}

func studentCreateStudentUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] student create-student -body JSON

学生を登録する。
    -body JSON: 

Example:
    %[1]s student create-student --body '{
      "address": "名古屋市中区三の丸三丁目1番2号",
      "date_of_birth": "2022-04-01T13:30:00+09:00",
      "expiration_date": "2027-03-31T00:00:00+09:00",
      "name": "鈴木太郎",
      "ruby": "スズキタロウ",
      "student_number": 12345
   }'
`, os.Args[0])
}
