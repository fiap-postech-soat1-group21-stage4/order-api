Feature: Create Order with Items

  Scenario: User creates an order with items
    Given an existing customer with ID "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
    And the order should have the following items:
      | ProductID                            | Quantity |
      | 550e8400-e29b-41d4-a716-446655440000 | 2        |
      | 6ba7b810-9dad-11d1-80b4-00c04fd430c8 | 3        |
    When the order is created
    Then the order should have the following items:
      | ProductID                            | Quantity |
      | 550e8400-e29b-41d4-a716-446655440000 | 2        |
      | 6ba7b810-9dad-11d1-80b4-00c04fd430c8 | 3        |