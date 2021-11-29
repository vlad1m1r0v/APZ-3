import Client from "./channels/client.js";

const client = new Client("http://localhost:8000");

// Scenario 1: Get disks
client.getDisks();

// Scenario 1: Get machines
client.getMachines();
