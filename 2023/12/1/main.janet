(defn chunk-by [f xs]
  (if (empty? xs)
    []
    (do
      (var pre nil)
      (def ret @[])
      (each x xs
        (def v (f x))
        (if (or (empty? ret) (not= v pre))
          (array/push ret @[x])
          (array/push (ret (- (length ret) 1)) x))
        (set pre v))
      ret)))

(defn parse-record [str]
  (def [x y] (string/split " " str))
  [(->> x
        string/bytes
        (map (fn [x] (case x 35 :broken 46 :working 63 :unknown)))
        (chunk-by identity)
        (map (fn [x] [(x 0) (length x)]))
        tuple/slice)
   (->> (string/split "," y)
        (map scan-number)
        tuple/slice)])

(defn add-unknown [as n]
  (match as
    [[:unknown num_unknown] & rest] [[:unknown (+ num_unknown n)] ;rest]
    (_ (> n 0)) [[:unknown n] ;as]
    _ as))

(defmacro memo-count-arrangements [conditions numbers memo]
  ~(do
     (def cached (get ,memo [,conditions ,numbers]))
     (if cached
       cached
       (let [value (count-arrangements ,conditions ,numbers ,memo)]
         (put ,memo [,conditions ,numbers] value)
         value))))

(defn count-arrangements [conditions numbers memo]
  (match [conditions numbers]
    [[[:broken num_broken1] [:broken num_broken2] & as] ns] (memo-count-arrangements [[:broken (+ num_broken1 num_broken2)] ;as] ns memo)
    [[[:working _] & as] ns] (memo-count-arrangements as ns memo)
    [[[:broken num_broken] [:unknown num_unknown] & as] [num_broken & ns]] (memo-count-arrangements (add-unknown as (- num_unknown 1)) ns memo)
    [[[:broken num_broken] & as] [num_broken & ns]] (memo-count-arrangements as ns memo)
    [[[:broken num_broken] [:unknown num_unknown] & as] [expected & ns]]
    (do
      (def missing (min (- expected num_broken) num_unknown))
      (if (>= missing 0)
        (memo-count-arrangements [[:broken (+ num_broken missing)] ;(add-unknown as (- num_unknown missing))] [expected ;ns] memo)
        0))
    [[[:unknown num_unknown] & as] ns]
    (let [as (add-unknown as (- num_unknown 1))
          as1 [[:working 1] ;as]
          as2 [[:broken 1] ;as]]
      (+ (memo-count-arrangements as1 ns memo)
         (memo-count-arrangements as2 ns memo)))
    ([_ _] (and (empty? conditions) (empty? numbers))) 1
    _ 0))

(def input (as-> (slurp "./input") _
                 (string/split "\n" _)
                 (array/slice _ 0 -2)
                 (map parse-record _)))

(->> input
     (map (fn [[conditions numbers]] (count-arrangements conditions numbers @{})))
     sum
     print)
