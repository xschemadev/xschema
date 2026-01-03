#!/usr/bin/env bun
import { createAdapterCLI } from "@xschema/core";
import { convert } from "./index";

createAdapterCLI(convert);
