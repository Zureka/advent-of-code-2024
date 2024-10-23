defmodule Day12Test do
  use ExUnit.Case

  @example_input """
  ???.### 1,1,3
  .??..??...?##. 1,1,3
  ?#?#?#?#?#?#?#? 1,3,1,6
  ????.#...#... 4,1,1
  ????.######..#####. 1,6,5
  ?###???????? 3,2,1
  """

  test "part 1 - determine possible permutations line 1" do
    input = "???.### 1,1,3"
    assert Day12.Part1.score_permutations(input) == 1
  end

  @tag :skip
  test "part 1 - determine possible permutations line 2" do
    input = ".??..??...?##. 1,1,3"
    assert Day12.Part1.score_permutations(input) == 4
  end

  @tag :skip
  test "part 1 - determine possible permutations line 3" do
    input = "?#?#?#?#?#?#?#? 1,3,1,6"
    assert Day12.Part1.solve(input) == 4
  end

  @tag :skip
  test "part 1 - determine possible permutations line 4" do
    input = ".??..??...?##. 1,1,3"
    assert Day12.Part1.solve(input) == 4
  end

  @tag :skip
  test "part 1 - determine possible permutations line 5" do
    input = ".??..??...?##. 1,1,3"
    assert Day12.Part1.solve(input) == 4
  end

  @tag :skip
  test "part 1 - determine possible permutations line 6" do
    input = ".??..??...?##. 1,1,3"
    assert Day12.Part1.solve(input) == 4
  end

  @tag :skip
  test "part 1 - calculates part 1 example correctly" do
    assert Day12.Part1.solve(@example_input) == 21
  end

  @tag :skip
  test "solves example input for part 2" do
    assert Day12.Part2.solve(@example_input) == 42
  end
end
