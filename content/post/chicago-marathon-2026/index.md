---
title: "Chicago Marathon 2026"
date: 2026-10-12
---

I ran the Chicago Marathon on October 11, 2026. These are my notes on race
preparation and execution (as usual, mostly written to remind myself about the
experience in decades to come).

{{<figure src="/images/post/boston-marathon-2023/finishline_small.jpg" link="/images/post/boston-marathon-2023/finishline.jpg">}}

## Result

2:53:54, 2187/54,000+ overall, 267/? M40-44

## Goals

* ✅ Goal A.1: Finish a successful training cycle.
* ✅ Goal A.2: Get to the starting line healthy.
* ❌ Stretch goal: 2:45 for Berlin
* ❌ Goal B: sub-2:48:42 for a new PB
* ❌ Goal C: sub-2:52 to beat NYC 2025

## Race Gap Analysis

Interactive pacing and gap comparison across all my marathons. By default, timing gaps are shown relative to Chicago 2026 (the zero baseline: positive values mean ahead/faster, negative values mean behind/slower). Use the controls to switch baselines, toggle distance units, or show/hide individual races.

{{< race-gap-chart >}}

## Weather Forecast Evolution

Ten-day evolution of the ECMWF and GFS ensemble model predictions leading up to race day on October 11, 2026 (7:00 AM – 3:00 PM CDT). Early model runs 7 to 10 days out projected warmer conditions approaching the mid-to-upper 60s (and mid-70s on GFS) with substantial ensemble spread. As race day approached, timing on a passing cold front sharpened, converging toward crisp running conditions in the mid-50s at the 7:00 AM start, warming to the low-60s by 10:00 AM and mid-70s by mid-afternoon.

Use the controls below to select between global models (**ECMWF** 50-member IFS vs. **GFS** 31-member GEFS), switch between metrics (temperature, dew point, wind speed, and rain probability), and scrub through race hours from 7 AM to 3 PM using the time slider against a static y-axis scale. Individual ensemble member curves and the operational deterministic runs are displayed alongside a horizontal dashed line marking the actual observed value on race day. Click the y-axis labels at any time to toggle between Imperial (°F, mph) and Metric (°C, km/h) units:

{{< weather-ensemble-chart >}}

### A Note on Model Grid Resolution and Ground Truth

{{< figure src="/images/post/chicago-marathon-2026/ecmwf-grid-cell-map.png" link="/images/post/chicago-marathon-2026/ecmwf-grid-cell-map.png" caption="The ECMWF 0.25° IFS grid cell (41.75°N–42.00°N, 87.50°W–87.75°W) centered over Chicago. Roughly half of the grid box covers the open waters of Lake Michigan, producing a coastal marine cool bias in numerical model reanalyses." >}}

The "observed" baseline in the chart above is taken from the official ASOS weather station at **Chicago Midway International Airport (KMDW)** (60°F at the 7:00 AM start, warming to 77°F by 2:00 PM and 3:00 PM), which best reflects the ambient urban conditions runners experienced along the race course.

When sampling global forecast models or numerical reanalyses (such as ECMWF ERA5) at the Grant Park race coordinates (41.8781°N, 87.6298°W), the simulated data represents a 0.25° grid cell (~21 km × 28 km). Because Grant Park sits directly on the shoreline, roughly 52% of this grid cell spans the open waters of Lake Michigan. In early autumn, the lake surface temperature hovers near 58°F (14°C). Numerical weather models compute area-averaged surface heat and moisture fluxes across the entire cell, blending the cool marine boundary layer into the 2-meter air temperature. This introduces a well-documented maritime cool bias of 3°F to 6°F compared to physical thermometers situated over inland urban pavement like Midway.

### Station Observations: Midway (KMDW) vs. O'Hare (KORD)

To examine how ambient temperatures and humidity evolved across the city on race day, the graph below compares simultaneous hourly records from Chicago's two primary ASOS stations—**Midway Airport (KMDW)** and **O'Hare Airport (KORD)**—between 7:00 AM and 3:00 PM:

{{< airport-obs-chart >}}

