Feature: Order Creation

  Scenario: Create order with valid input
    Given the following order details
      | CustomerID                             | ProductID                               | Quantity |
      | 6ba7b810-9dad-11d1-80b4-00c04fd430c8   | 550e8400-e29b-41d4-a716-446655440000    | 2        |
      | 6ba7b810-9dad-11d1-80b4-00c04fd430c8   | 6ba7b810-9dad-11d1-80b4-00c04fd430c8    | 1        |
    When a request is made to create the order
    Then the response should have status code 201
    And the response body should match the expected order details
