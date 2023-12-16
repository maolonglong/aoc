#!/usr/bin/env janet

(defn next-dirs
  [byte dir]
  (match [(string/from-bytes byte) dir]
    ["." _] [dir]
    ["/" [-1 0]] [[0 1]]
    ["/" [1 0]] [[0 -1]]
    ["/" [0 -1]] [[1 0]]
    ["/" [0 1]] [[-1 0]]
    ["\\" [-1 0]] [[0 -1]]
    ["\\" [1 0]] [[0 1]]
    ["\\" [0 -1]] [[-1 0]]
    ["\\" [0 1]] [[1 0]]
    ["|" [_ 0]] [dir]
    ["|" [0 _]] [[-1 0] [1 0]]
    ["-" [_ 0]] [[0 -1] [0 1]]
    ["-" [0 _]] [dir]))

(defn grid-get
  [grid row col]
  (-> grid
      (get row)
      (get col)))

(defn pos-add
  [[row col] [i j]]
  [(+ row i) (+ col j)])

(defn move-beam
  [grid pos dir]
  (def memo @{})
  (def energized @{})
  (defn helper
    [pos dir]
    (when (nil? (memo [pos dir]))
      (def [row col] pos)
      (put memo [pos dir] true)
      (when-let [byte (grid-get grid row col)]
        (put energized pos true)
        (def dirs (next-dirs byte dir))
        (each dir dirs
          (helper (pos-add pos dir) dir)))))
  (helper pos dir)
  energized)

(def grid
  (->> "./input"
       slurp
       string/trim
       (string/split "\n")))

(-> grid
    (move-beam [0 0] [0 1])
    length
    print)
