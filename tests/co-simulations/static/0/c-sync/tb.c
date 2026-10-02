#include <assert.h>
#include <stdio.h>

#include "cosim_iface.h"

#include "amba_fbd.h"
#include "main.h"
#define AMBA_FBD_IFACE &iface


int main(int argc, char *argv[]) {
	assert(argc == 3);

	amba_fbd_iface_t iface = cosim_iface_iface();

	cosim_iface_init(argv[1], argv[2], NULL);

	uint32_t id;
	amba_fbd_read(main_ID, &id);
	if (id != amba_fbd_main_ID) {
		fprintf(stderr, "read wrong ID %x, expecting %x\n", id, amba_fbd_main_ID);
		cosim_iface_end(1);
	}

	uint32_t timestamp;
	amba_fbd_read(main_TIMESTAMP, &timestamp);
	if (timestamp != amba_fbd_main_TIMESTAMP) {
		fprintf(stderr, "read wrong TIMESTAMP %x, expecting %x\n", id, amba_fbd_main_TIMESTAMP);
		cosim_iface_end(1);
	}

	cosim_iface_end(0);

	return 0;
}
