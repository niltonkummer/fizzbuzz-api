Feature: Stats API
  In order to use the Stats API
  As a user
  I need to be able to send requests and receive responses
  Scenario: then user try to get stats
    When I send "GET" request to "/stats":
    Then the response code should be 200
    And the response payload should match json:
        """
         {
         "int1": 3,
         "int2": 5,
         "limit": 15,
         "str1": "Fizz",
            "str2": "Buzz",
            "hits": 1
         }
        """
  Scenario: then user try to get stats with zero hits (reset stats)
    When I send "GET" request to "/stats":
    Then the response code should be 404
    And the response payload should match json:
        """
         {
         "code": "no_requests_found",
         "message":"No requests found in the statistics"
         }
        """