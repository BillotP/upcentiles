# Upcentiles , an analytics api for Upfluence SSE stream datas

**🧑‍🔬 Still a work in progress**

## 🥅 Goal

- For each http api request on `/analysis?duration=x&dimension=y`,

- Fetch social media post events from Upfluence SSE api during **x** duration

- Parse the received datas to extract **timestamp** and **y** dimension values

- Get the first and last **timestamp** values of the received datas

- Sort the datas by **y** dimension values

- Compute 3 percentiles values from the sorted datas : p_50 (median), p_90 and p_99

- Return a JSON object with total number of items in set, first and last timestamps and the computed percentiles.

## 📓 Methodology

- [x] Read the goals (the subject)

- [x] Upstream API discovery via cli tools (of course it's **curl**)

- [x] Write down data objects models

- [x] Implement a draft (quick & dirty) project version with a minimal separation of concerns

- [ ] Re-Read the goals to avoid missing features

- [ ] Write unit tests with mocked upstream api for critical components (data fetching and statistics computation)

- [ ] Refactoring to stay DRY and best-practices compliant

- [ ] Repeat

## 📓 Implementation

Current implementation consist of :

- A Golang api server with [`echo`](https://pkg.go.dev/github.com/labstack/echo/v4) framework

- A package for api structs definitions

- A listen the stream and send the events to a **chan** during x duration method,

- A sort set and compute centiles method

- A GET api handler

- A main function to start the api

## 💥 Limits

- Memory usage++

- Each client connection to the api create a new upstream api listener, the connection is not shared

- Current draft implementation parse and format all the upstream api datas, I guess we should avoid to save all the keys we dont need.

- The error handling in [Subscribe](./internal/upfluence/subscribe.go) method is a bit too strict and doesnt allow for a retry if connection failed or is interrupted

- All the running env (eg the PORT this server is listening on or the upstream api URL) is hardcoded, this is bad practice

- The [Percentiles](./internal/stats/stats.go) method shouldn't have any reference to the api models and should only take the unsorted slice of "dimension" values

## 💥 Misc

- Wondering if "NonAccurate" events should be removed from the statistical set (for example, if _likes_ is the dimension, _Article_ type of social event does not have this property)NonAccurate

- Stat method and api models could have been part of the public api (`pkg` directory) for this project but as its only a test, all this parts have been put in `internal` directory

- Slice sorting have been implemented using go slices [Sort](https://pkg.go.dev/slices#Sort) package, as no external lib is allowed it is planned to re-implement it with either a partition sort for small list and quicksort for largest ones
