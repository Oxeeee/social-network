import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

using Http;
using OpenAPI;

@service(#{ title: "Connectify" })
@server("https://localhost:3000", "Localhost")
@info(#{ version: "0.0.0" })
namespace Connectify;

model LoginPayload {
  email: string;
  @minLength(2) password: string;
}

model LoginResponse {
  accessToken: string;
}

model RegisterPayload {
  email: string;
  @minLength(8) password: string;
  @minLength(2) name: string;
  @minLength(2) surname: string;
  @minLength(2) username: string;
}

model User {
  email: string;
  name: string;
  surname: string;
  username: string;
  description: string;
  photo: string;
}

model UserUpdate {
  email?: string;
  name?: string;
  surname?: string;
  username?: string;
  description?: string;
  photo?: string;
}

@tag("Auth")
@route("/login")
@post
op LoginUser_loginUser(@body body: LoginPayload): Body<{
  message: string;
  error: string;
  data: LoginResponse;
}>;

@tag("Auth")
@route("/register")
@post
op RegisterUser_loginUser(@body body: RegisterPayload): CreatedResponse;

@tag("User")
@route("/user/{userId}")
@patch
op UserOperations_editUser(@path userId: string, @body body: UserUpdate): Body<{
  message: string;
  error: string;
  data: User;
}>;
