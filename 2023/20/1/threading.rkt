#lang racket

(provide (all-defined-out))

(define-syntax ->
  (syntax-rules ()
    [(_ expr) expr]
    [(_ expr (proc args ...) rest ...)
     (-> (proc expr args ...) rest ...)]
    [(_ expr proc rest ...) (-> (proc expr) rest ...)]))

(define-syntax ->>
  (syntax-rules ()
    [(_ expr) expr]
    [(_ expr (proc args ...) rest ...)
     (->> (proc args ... expr) rest ...)]
    [(_ expr proc rest ...) (->> (proc expr) rest ...)]))

(define-syntax as->
  (lambda (stx)
    (syntax-case stx ()
      [(_ expr id) (identifier? #'id) #'expr]
      [(_ expr id (proc args ...) rest ...)
       (identifier? #'id)
       #'(as-> (let ([id expr]) (proc args ...)) id rest ...)]
      [(_ expr id proc rest ...)
       (identifier? #'id)
       #'(as-> (let ([id expr]) (proc)) id rest ...)])))
