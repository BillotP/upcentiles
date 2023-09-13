# Upcentiles , an analytics api for Upfluence SSE stream datas

## 🥅 Goals \*

- For each http api request on `/analysis?duration=x&dimension=y`,

- Fetch social media post events from Upfluence SSE api during **x** duration

- Parse the received datas to extract **timestamp** and **y** dimension values

- Get the first and last **timestamp** values of the received datas

- Sort the datas by **y** dimension values

- Compute 3 percentiles values from the sorted datas : p_50 (median), p_90 and p_99

- Return a JSON object with total number of items in set, first and last timestamps and the computed percentiles.

\* Integrality of the goals /subject of this exercise is kept private

## 📓 Methodology

- [x] Read the goals (the subject)

- [x] Upstream API discovery via cli tools (of course it's **curl**)

- [x] Write down data objects models

- [x] Implement a draft (quick & dirty) project version with a minimal separation of concerns

- [x] Re-Read the goals to avoid missing features

- [x] Write unit tests with mocked upstream api for critical components (parameter validation and statistics computation)

- [x] Refactoring to stay DRY and best-practices compliant

## 📓 Implementation

Current implementation consist of :

- A Golang api server with [`echo`](https://pkg.go.dev/github.com/labstack/echo/v4) framework

- A package for api structs definitions in [api](./internal/api/api.go)

- A listen the stream and send the events to a **chan** during x duration method ([Subscribe](./internal/upfluence/subscribe.go)),

- A [sort set](./internal/stats/sort.go) and [compute centiles](./internal/stats/stats.go) methods

- A GET [api handler](./internal/handler/analysis.go)

- A main function to start the api server

## 💥 Limits

- Memory usage is a bit too high in [SubscribeFull](./internal/upfluence/subscribe.go#L19) as we deserialize and save the whole [StreamEvent](./internal/upfluence/models.go#L17) in ram without the need for the extra fields it contains. This is a bit mitigated in [SubscribeLight](./internal/upfluence/subscribe.go#L74) method as only the timestamps and dimension values are sent to the chan.

- Each client connection to the api create a new upstream api listener, the connection is not shared, I could have used a shared stream of event with mutexes for read for every consumer request.

- ~~Current draft implementation parse and format all the upstream api datas, I guess we should avoid to save all the keys we dont need.~~ (cf Memory Usage).

- The error handling in [Subscribe](./internal/upfluence/subscribe.go) methods is a bit too strict and doesnt allow for a retry if connection failed or is interrupted.

- All the running env (eg the PORT this server is listening on or the upstream api URL) is hardcoded, this is bad practice.

- This project is using the golang "log" package to print info, warnings and errors to stdout , the message format and the logger package should reflect the one used by the organisation.

- ~~The [Percentiles](./internal/stats/stats.go) method shouldn't have any reference to the api models and should only take the unsorted slice of "dimension" values for a better separation of concerns~~ (Fixed in [PercentilesV2](./internal/stats/stats.go#L38)).

## 🦄 Misc

- Wondering if "NonAccurate" events should be removed from the statistical set (for example, if _likes_ is the dimension,[ _Article_](./internal/upfluence/models.go#L77) type of stream event does not have this property).

- Wondering if some event kinds property without the same name but representing the same dimension should be accounted as so (eg : Is _repins_ the same as _retweets_).

- Stat method and api models could have been part of the public api (`pkg` directory) for this project but as its only a test, all this parts have been put in `internal` directory.

- Regarding the [upfluence](./internal/upfluence/) package, the organisation should have a dedicated client with proper structs and options, this implementation is volontary simple.

- ~~Slice sorting have been implemented using go slices [Sort](https://pkg.go.dev/slices#Sort) package, as no external lib is allowed it is planned to re-implement it with either a partition sort for small list and quicksort for largest ones~~ ( EDIT : A naive Quicksort and Partition sort algo is implemented on the [sort package](./internal/stats/sort.go))

- In Subscribe method, a strings.Replace method is used to isolate the json string from the `data: ` prefix, seems a bit overkilled.

## API Documentation

<code>GET</code> <b>/analysis</b></code>
: Getting statistical analysis for Upfluence SSE stream datas

##### Parameters

> | name      | type     | data type                | description                                                 |
> | --------- | -------- | ------------------------ | ----------------------------------------------------------- |
> | duration  | required | string (query parameter) | A duration eg: 10s, 24h, 1d ...                             |
> | dimension | required | string (query parameter) | An analyis dimension : likes, comments, favorites, retweets |

##### Responses

> | http code | content-type       | response                                                                                                                                                               |
> | --------- | ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
> | `200`     | `application/json` | `{"total_posts": number,"minimum_timestamp":UNIX timestamp number,"maximum_timestamp":UNIX timestamp number,"likes_p50":number,"likes_p90":number,"likes_p99":number}` |
> | `400`     | `application/json` | `{"error":"a usefull error message describing the bad parameter"}`                                                                                                     |
> | `204`     | NA                 | No content if no SSE data match the required **dimension** parameter                                                                                                   |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:8080/analysis?duration=10s&dimension=likes
> ```

## Usage

- Install a recent Go / Golang version in your favorite distribution ([HELP](https://go.dev/doc/install)).

- Optionally install GNU `make` tool and a container building tool compactible with the `dockerfile` recipe (eg: buildah, podman, docker cli ..).

- Once done, you are ready to **run the api** command with `go run cmd/server/main.go` executed on your favorite shell interface.

- If you have the `make` tool setup, run `make help` or just `make` from the repo root to see the available commands.