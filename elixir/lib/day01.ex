defmodule Day01.Part1 do
  def solve(input) do
    [left, right] = createSublists(input)

    Enum.zip(left, right)
    |> Enum.reduce(0, fn {a, b}, acc -> acc + abs(b - a) end)
  end

  defp createSublists(input) do
    [left, right] =
      input
      |> String.split("\n", trim: true)
      |> Enum.reduce([[], []], fn line, [left, right] ->
        [a, b] =
          String.split(line, "   ", trim: true)
          |> Enum.map(&String.to_integer/1)

        [[a | left], [b | right]]
      end)

    [Enum.sort(left), Enum.sort(right)]
  end
end

defmodule Day01.Part2 do
  def solve(input) do
    [left, right] = createSublists(input)

    left
    |> Enum.reduce(0, fn x, total ->
      count =
        Enum.reduce(right, 0, fn y, acc ->
          if x == y, do: acc + 1, else: acc
        end)

      total + x * count
    end)
  end

  defp createSublists(input) do
    [left, right] =
      input
      |> String.split("\n", trim: true)
      |> Enum.reduce([[], []], fn line, [left, right] ->
        [a, b] =
          String.split(line, "   ", trim: true)
          |> Enum.map(&String.to_integer/1)

        [[a | left], [b | right]]
      end)

    [Enum.sort(left), Enum.sort(right)]
  end
end

defmodule Mix.Tasks.Day01 do
  use Mix.Task

  def run(_) do
    {:ok, input} = File.read("inputs/day01-input.txt")

    IO.puts("--- Part 1 ---")
    IO.inspect(Day01.Part1.solve(input))
    IO.puts("")
    IO.puts("--- Part 2 ---")
    IO.puts(Day01.Part2.solve(input))
  end
end
