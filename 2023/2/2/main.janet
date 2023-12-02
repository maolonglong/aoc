#!/usr/bin/env janet

(defmacro max= [x & ns] ~(set ,x (,max ,x ,;ns)))

(def game-peg
  ~{:main (* (any :game) -1)
    :game (group
            {:main (* :s* :begin :sets :s*)
             # Game 1:
             :begin (* "Game" :s+ (cmt (<- :d+) ,scan-number) ":" :s+)
             # 3 blue
             :cubes (group (* (cmt (<- :d+) ,scan-number) :s+ (<- :w+)))
             # 3 blue, 4 red
             :set (group (* :cubes (any (* "," :s+ :cubes))))
             # 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green ...
             :sets (group (* :set (any (* ";" :s+ :set))))})})

(def input (string/trim (slurp "./input")))
(def games (peg/match game-peg input))

(defn game->cost [game]
  (def [id sets] game)
  (var cost @[0 0 0])
  (each s sets
    (each [cnt color] s
      (case color
        "red" (max= (cost 0) cnt)
        "green" (max= (cost 1) cnt)
        "blue" (max= (cost 2) cnt))))
  (tuple/slice cost))

(defn power [game]
  (def [r g b] (game->cost game))
  (* r g b))

(->> games
     (map power)
     (reduce + 0)
     (print))
