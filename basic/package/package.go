/*
커스텀 패키지

$GOPATH/src/my_package
$GOPATH/src/my_package/pkg1
$GOPATH/src/my_package/pkg2

Go 코드 컨벤션 패키지 이름
패키지와 디렉터리는 같은 이름으로 생성
파일 한개로 구성된 패키지는 소스파일의 이름을 패키지명.go로 생성

main package를 작성후 go build 명령은 해당 패키지 디렉터리에 실행 파일 생성
go install 명령은 bin 디렉터리에 실행 파일 생성

서드 파티 패키지
go get ~

*/

/* 테스트
go test
테스트 파일 : 파일 이름이 _test.go로 끝나는 파일
테스트 함수 : 함수 이름이 Test로 시작하며 *testing.T를 매개변수로 받는 함수

테스트 함수의 성공/실패는 테스트 실패 메서드를 호출하는 것으로 결정
ㅡ testing.T.Errorf()
ㅡ testing.T.Fail()
ㅡ testing.t.Fatal()
*/

package main
