#!/usr/bin/env python

# Imports
from twisted.internet import reactor
import time
import appdash

# Appdash: Twisted Collector
from appdash.twcollector import RemoteCollectorFactory

# Create a remote appdash collector.
collector = RemoteCollectorFactory(reactor, debug=True)

# Create a trace.
trace = appdash.SpanID(root=True)

# Generate a few spans with some annotations.
span = trace
for i in range(0, 7):
    # Name the span.
    if i == 0:
        collector.collect(span, *appdash.MarshalEvent(appdash.SpanNameEvent("Request")))
    else:
        collector.collect(span, *appdash.MarshalEvent(appdash.SpanNameEvent("SQL Query")))

    # Marshal some events into annotations and collect them.
    sendTime = time.time()
    recvTime = sendTime + 2
    collector.collect(span, *appdash.MarshalEvent(appdash.SQLEvent(
        "SELECT * FROM table_name;",
        sendTime,
        recv = recvTime, # optional: default is current time.
        tag = "foobar",  # optional: user-specific tag, useful for e.g. filtering.
    )))

    if i % 2 == 0:
        collector.collect(span, *appdash.MarshalEvent(appdash.LogEvent("Hello world!")))

    # Create a new child span whose parent is the last span we created.
    span = appdash.SpanID(parent=span)

# Have Twisted perform the connection and run.
reactor.connectTCP("", 7701, collector)
reactor.run()
