# Go

## 코드 구조

> Workspace > Repository > packages > Go file

root 디렉토리에 src, pkg, bin 디렉토리가 있으면 workspace.

go tool을 통해서 src의 소스를 빌드해서 pkg와 bin에 생성함

- src -> 소스
- pkg -> 라이브러리?
- bin -> 실행 파일

한 패키지 내의 모든 파일은 동일한 패키지 이름을 써야함.

import path의 마지막 이름을 주로 사용한다고 함
> "crypto/rot13" -> rot13

Executable일 경우 **반드시** main package

### 독특한 Package 시스템

package가 곧 저장소의 url이 됨

> ex) "https://github.com/soldier443/hello"에 저장소가 있으면

go get github.com/soldier4443/hello로 원격 소스를 받아서 install할 수 있음

-> git clone과 유사하고 간편?

-> 다른 go package들을 아주 쉽게 가져다 쓸 수 있게 됨!

## 명령어

- go build - build package

- go install - build package and generate output file. Also install other dependencies.

- go get - download remote package and build? or not

## Private / Public

first letter, 첫 글자에 따라서 private/public 결정.

print - private
Print - public

## Pointer

- &: 주소
- *: 주소에 있는 값

## Struct

연관된 값들을 모아놓을 수 있음

type ~~~ struct {
  a unit16
  b float32
  ...
}

## 메소드

Method Receiver - 클래스가 없는 대신, 포인터와 인터페이스를 제외한 모든 타입에 대해 메서드를 추가할 수 있다.

형태는 다음과 같다.

```go
func (v T) funcName(param U) V {
  // Implementation
}
```

### Value Receiver

타입을 복사해서 전달받는다. 이 때문에 어떤 변경을 가하더라도 해당 메서드 내에서만 유지된다.

값을 수정할 수 없기 때문에 예상치 못하게 값이 변경되어 side-effect를 야기하는 문제를 피할 수 있을 것 같다.

```go
type Point struct {
  x float32
  y float32
}

func (p Point) negate() Point {
  return {-p.x, -p.y}
}
```

### Pointer Recevier

타입의 포인터를 전달받는다. 타입 자체에 수정이 필요한 경우 반드시 Pointer Receiver를 사용해야 한다.

포인터를 전달받기 때문에 타입의 크기가 클 때(거대한 구조체 등) 유용하기도 하다.

```go
func (p *Point) multiply(multiplier float32) Point {
  p.x *= multiplier
  p.y *= multiplier
}
```

## Interface

Go에도 인터페이스가 있다구??
그렇다. 있다!

다른 언어들과 비슷하게 인터페이스는 해당 구조체?의 동작을 **정의**한다.
그런데 그 연결이 조금 느슨하다느 느낌을 받음.

```go
type Vector2 struct {
  x float64
  y float64
}

type Vector3 struct {
  x float64
  y float64
  z float64
}

type Measurable interface {
  magnitude() float64
}

func (v Vector2) magnitude() float64 {
  return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v Vector3) magnitude() float64 {
  return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func main() {
  vec2 := Vector2{10.0, 15.0}
  vec3 := Vector3{x: 4.0, y: 11.0}

  fmt.Println(vec2, "->", Measurable(vec2).magnitude())
  fmt.Println(vec3, "->", Measurable(vec3).magnitude())
}
```

메서드 자체가 구조체 내부에 있지 않고 외부에 있는데
"어떤 struct가 어떤 interface를 구현하고 있다/아니다"를 어떻게 판단하는지 궁금해졌다.

[Interface에 대한 자료](https://flaviocopes.com/go-empty-interface/)를 발견하고 궁금증이 풀렸다.
해당 글에서는 **Duck typing**의 개념을 이야기하고 있다.

> If it walks like duck, swims like a duck and quacks like a duck, it’s a duck

무언가가 오리처럼 걷고, 오리처럼 운다면 그것은 오리라는 것.
그래서 Go에는 **implements** 키워드같이 명시적인 인터페이스 구현에 대한 선언이 존재하지 않는다.

그렇기 때문에 **인터페이스를 정의한 패키지로부터 구현 패키지를 분리**해준다고 하긴 하는데, 아직은 잘 모르겠다..ㅠㅠ

### 인터페이스 + 에러 처리

Go는 에러를 처리할 때 error 인터페이스를 사용한다.

```go
type error interface {
  Error() string
}
```

## 여담

정적 타입 언어
-> 타입이 달라지지 않음

import를 할 때는 "", 여러 개를 할 때는 () 안에 "",
서브 패키지는 /로 구분

ex) "fmt", "math/rand"

main 패키지가 아닌 것은 실행할 수 없다.

전체적으로 C+++++++++같은 느낌?

Composite literal에서는 마지막 라인에도 ,가 있어야 함..?

Type Conversion / Type Assertion의 차이??

Go 자체에서 지원하는 유용한 패키지들이 정말 많다.
한 번 알아보면 좋을듯.
