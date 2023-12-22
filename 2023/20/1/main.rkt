#lang racket

(require "threading.rkt"
         data/queue)

(define (parse)
  (define lines (with-input-from-file "./input"
                  (lambda ()
                    (let loop ([line (read-line)]
                               [lines null])
                      (if (eof-object? line)
                          (reverse lines)
                          (loop (read-line) (cons line lines)))))))
  (map (lambda (line)
         (match-define (list id signals) (string-split line " -> "))
         (define-values (type name)
           (cond
             [(eqv? (string-ref id 0) #\%) (values 'flip-flop (substring id 1))]
             [(eqv? (string-ref id 0) #\&) (values 'conjunction (substring id 1))]
             [else (values 'broadcast "broadcast")]))
         (list type name (string-split signals ", ")))
       lines))

(define (prepare-modules input)
  (define sources (->> input
                       (append-map
                        (match-lambda [(list _ name destinations)
                                       (map (lambda (dest) (list dest name)) destinations)]))
                       (group-by car)
                       (map (lambda (xs)
                              (define key (caar xs))
                              (define value (map (lambda (x) (cadr x)) xs))
                              (cons key value)))
                       (make-immutable-hash)))
  (->> input
       (map (match-lambda [(list type name destinations)
                           (case type
                             ['flip-flop (list name type destinations #f)]
                             ['conjunction
                              (define state (->> (hash-ref sources name)
                                                 (map (lambda (x) (list x #f)))))
                              (list name type destinations state)]
                             ['broadcast (list name type destinations #f)])]))
       make-immutable-hash))

(define (signal modules item)
  (match-define (list name from value) item)
  (define default (lambda () (list 'output null #f)))
  (match-define (list type destinations state) (hash-ref modules name default))
  (case type
    ['flip-flop
     (cond
       [value (list modules null)]
       [else
        (define new-state (not state))
        (define new-modules (hash-set modules name (list type destinations new-state)))
        (list new-modules (map (lambda (x) (list x name new-state)) destinations))])]
    ['conjunction
     (define new-state (map
                        (match-lambda [(list state-name state-value)
                                       (if (equal? state-name from)
                                           (list state-name value)
                                           (list state-name state-value))])
                        state))
     (define new-value (not (null? (filter-not cadr new-state))))
     (define new-modules (hash-set modules name (list type destinations new-state)))
     (list new-modules (map (lambda (x) (list x name new-value)) destinations))]
    ['broadcast
     (list modules (map (lambda (x) (list x name value)) destinations))]
    ['output (list modules null)]))

(define (pulse modules)
  (define q (make-queue))
  (enqueue! q (list "broadcast" "broadcast" #f))
  (let loop ([low 0]
             [high 0]
             [modules modules])
    (cond
      [(queue-empty? q) (list (list low high) modules)]
      [else
       (define item (dequeue! q))
       (match-define (list new-modules items) (signal modules item))
       (for-each (lambda (item) (enqueue! q item)) items)
       (match item
         [(list _ _ #f)
          (loop (add1 low) high new-modules)]
         [(list _ _ #t)
          (loop low (add1 high) new-modules)])])))

(define input (parse))
(define modules (prepare-modules input))

(let loop ([i 0]
           [modules modules]
           [low-acc 0]
           [high-acc 0])
  (cond
    [(= i 1000) (* low-acc high-acc)]
    [else
     (match-define (list (list low high) new-modules) (pulse modules))
     (loop (add1 i) new-modules (+ low-acc low) (+ high-acc high))]))
