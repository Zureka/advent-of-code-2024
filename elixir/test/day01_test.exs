defmodule Day01Test do
  use ExUnit.Case

  @example_input """
  3   4
  4   3
  2   5
  1   3
  3   9
  3   3
  """

  test "solves example input for part 1" do
    assert Day01.Part1.solve(@example_input) == 11
  end

  test "solves example input for part 2" do
    assert Day01.Part2.solve(@example_input) == 31
  end
end
