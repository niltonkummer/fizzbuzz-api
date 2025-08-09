Feature: FizzBuzz API
    In order to use the FizzBuzz API
    As a user
    I need to be able to send requests and receive responses

    Scenario: then user try to create a fizzbuzz
        When I send "POST" request to "/fizzbuzz" with payload:
        """
        {
            "int1": 3,
            "int2": 5,
            "limit": 15,
            "str1": "Fizz",
            "str2": "Buzz"
        }
        """
        Then the response code should be 200
        And the response payload should match json:
        """
          {"response": "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"}
        """
  Scenario: then user try to create a fizzbuzz with invalid parameters
    When I send "POST" request to "/fizzbuzz" with payload:
        """
        {
            "int1": 0,
            "int2": 1,
            "limit": 15,
            "str1": "Fizz",
            "str2": "Buzz"
        }
        """
    Then the response code should be 400
    And the response payload should match json:
        """
        {
            "code": "invalid_request",
            "message": "int1 must be greater than 1"
        }
        """