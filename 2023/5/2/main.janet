#!/usr/bin/env janet

(defn chunk [n xs]
  (if (empty? xs)
    []
    (let [head (take n xs)
          rest (drop n xs)]
      [head ;(chunk n rest)])))

(defn seeds-part2 [xs] (chunk 2 xs))

(def input-peg
  ~{:main (* :seeds :maps -1)
    :number (cmt (<- :d+) ,scan-number)
    :seeds (cmt
             (* "seeds:" (group (some (* :s* :number))))
             ,seeds-part2)
    :maps (group (some (* :s* (group :map))))
    :map (* :map-name :s+ "map:\n" (some (group :map-line)))
    :map-name (some (+ :w "-"))
    :map-line (* (some (* (any " ") :number)) "\n")})

(defn get-range-overlap [(start1 n1) (start2 n2)]
  (def end1 (+ start1 n1))
  (def end2 (+ start2 n2))
  (unless (or (<= end1 start2) (<= end2 start1))
    (let [start (max start1 start2)
          end (min end1 end2)
          n (- end start)]
      [start n])))

(defn remove-overlap [(start1 n1) (start2 n2)]
  (def end1 (+ start1 n1))
  (def end2 (+ start2 n2))
  (def ret @[])
  (if (and (> end1 start2) (< start1 start2))
    (array/push ret [start1 (- start2 start1)]))
  (if (and (> end2 start1) (< end2 end1))
    (array/push ret [end2 (- end1 end2)]))
  ret)

(defn source->target [source m]
  (mapcat (fn [s]
            (var ret [s])
            (each [t-start s-start n] m
              (def overlap (get-range-overlap s [s-start n]))
              (unless (nil? overlap)
                (def [start n] overlap)
                (def offset (- t-start s-start))
                (set ret [[(+ start offset) n]
                          ;(source->target (remove-overlap s overlap) m)])
                (break)))
            ret)
          source))

(def input (->> (slurp "./input")
                (peg/match input-peg)))

(def [seeds maps] input)
(->> (reduce source->target seeds maps)
     (map 0)
     min-of
     print)
