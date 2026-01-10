#!/usr/bin/env node
import { createAdapterCLI } from "@xschemadev/core";

import { convert } from "./index.js";

createAdapterCLI(convert);
