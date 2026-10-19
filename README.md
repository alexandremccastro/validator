# Go Validate
A struct validation lib for GoLang

## Test
To execute tests in this library, you need to execute the following command:

```sh
make test
```

## Coverage
To check coverage you can execute the following command:
```sh
make test-cov
```


## How to use?
This library uses the native lang support for tags in struct fields to empower validation, as you can see below:

```golang
type LoginRequest struct {
  Email string `json:"email" validate:"email,required"`
  Password string `json:"password" validate:"required"`
}
```

To execute the validation and the the decoded data or errors execute the `DecodeAndValidate` function:

```golang
credentials, errors := validator.DecodeAndValidate[LoginRequest](req.Body)
```


### Available Rules

| Rule | Description |
| -------- | -------- |
| email | Check if the field is a valid email |
| required | Ensure the value is set |
| trim | Trim strings in both left and right to avoid whitespaces in the edges |
| min | Ensure that the field value match a minimum value |
| max | Ensure the field value match a maximum value |
| in | Ensure the field value is match on a predefined list |
| len | Ensure the value len match a specific value |
