# Evaluate Reverse Polish Notation

> NeetCode 150 · Stack — operands push, operators pop two and push one; postfix needs no precedence, no parentheses, no parser. Go, stdlib only, tested.

## Problem

Evaluate an arithmetic expression in postfix form (`2 1 + 3 *` = 9) with +, −, ×, and division truncating toward zero.

## Approach

RPN exists *because* a stack evaluates it trivially — operator precedence was resolved by whoever produced the postfix order, so evaluation is a single pass: numbers push; an operator pops its two operands (note the order: second pop is the *left* operand) and pushes the result. One value must remain at the end. Go's integer division already truncates toward zero (unlike Python's floor), so the spec's division rule costs nothing — the `6 / -132 = 0` test documents it. Malformed inputs (operator starvation, leftovers, unknown tokens, ÷0) return ok=false rather than panicking — the error cases are most of the real implementation.

**Complexity:** O(n) time, O(n) stack.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
