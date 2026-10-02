#!/bin/bash
set -e

CONTEXT=$1
ENTITY=$2
HDL=$3

LOGDIR="/tmp/amba-fbd/$CONTEXT/$ENTITY/python"
PYTHON_FILE="../../../tests/co-simulations/$CONTEXT/$ENTITY/python/tb.py"

export PYTHONPATH="$PYTHONPATH:$PWD/../../../tests/co-simulations/common/python/"
export PYTHONPATH="$PYTHONPATH:$PWD/../../amba-fbd/$CONTEXT/$ENTITY/python/$HDL"

mkdir -p $LOGDIR
if [ ! -f "$PYTHON_FILE" ]; then
	>&2 echo "$PYTHON_FILE not found"
	exit 1
fi
python3 $PYTHON_FILE \
	/tmp/amba-fbd/$CONTEXT/$ENTITY/python_"$HDL" \
	/tmp/amba-fbd/$CONTEXT/$ENTITY/"$HDL"_python \
	> "$LOGDIR/$HDL.log" 2>&1 &
