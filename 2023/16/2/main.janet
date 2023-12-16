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

(def row-max (- (length grid) 1))
(def col-max (- (length (grid 0)) 1))

(def start-beams
  (array/concat
    (mapcat (fn [col] [[[0 col] [1 0]] [[row-max col] [-1 0]]]) (range col-max))
    (mapcat (fn [row] [[[row 0] [0 1]] [[row col-max] [0 -1]]]) (range row-max))))

(->> (map (fn [[pos dir]] (move-beam grid pos dir)) start-beams)
     (map length)
     max-of
     pp)
