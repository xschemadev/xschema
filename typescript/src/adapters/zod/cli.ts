#!/usr/bin/env node
import { createAdapterCLI } from "../../index.js";
import { convert } from "./index.js";

createAdapterCLI(convert);
