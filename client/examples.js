import Client from "./requests/client.js";

const client = new Client("http://localhost:8000");

(async function () {
	// Scenario 1: Connect disk to machine
	console.log(await client.connectToMachine(1, 1));
	// Successful disk request

	// Scenario 2: Get machines
	console.log(await client.getMachines());
	// [
	//     { id: 2, name: 'server-2', cpuCount: '8', totalDiskSpace: '0' },
	//     { id: 3, name: 'server-3', cpuCount: '2', totalDiskSpace: '256' },
	//     { id: 1, name: 'server-1', cpuCount: '4', totalDiskSpace: '1024' }
	//   ]
})();
