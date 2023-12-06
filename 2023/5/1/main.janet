#!/usr/bin/env janet

(def input-peg
  ~{:main (* :seeds :maps -1)
    :number (cmt (<- :d+) ,scan-number)
    :seeds (* "seeds:" (group (some (* :s* :number))))
    :maps (group (some (* :s* (group :map))))
    :map (* :map-name :s+ "map:\n" (some (group :map-line)))
    :map-name (some (+ :w "-"))
    :map-line (* (some (* (any " ") :number)) "\n")})

(defn source->target [source m]
  (map (fn [s]
         (var ret s)
         (each [t-start s-start n] m
           (when (and (>= s s-start)
                      (< s (+ s-start n)))
             (set ret (+ t-start (- s s-start)))
             (break)))
         ret)
       source))

(def input (->> (slurp "./input")
                (peg/match input-peg)))

(def [seeds maps] input)
(->> (reduce source->target seeds maps)
     min-of
     print)
