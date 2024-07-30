#!/bin/bash

flutter build web --web-renderer canvaskit --release \
&& cp -R build/web/* ../public/
