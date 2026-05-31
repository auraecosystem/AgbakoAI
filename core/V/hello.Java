HttpResponse<String> response = Unirest.get("https://api.dev.Aexample.com/core/v1/hello")
  .header("Authorization", "Basic <encoded-value>")
  .asString();
