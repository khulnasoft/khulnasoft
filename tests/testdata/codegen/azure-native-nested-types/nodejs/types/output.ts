// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace documentdb {
    export interface CompositePathResponse {
        /**
         * Sort order for composite paths.
         */
        order?: string;
        /**
         * The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)
         */
        path?: string;
    }

    /**
     * Cosmos DB indexing policy
     */
    export interface IndexingPolicyResponse {
        /**
         * List of composite path list
         */
        compositeIndexes?: outputs.documentdb.CompositePathResponse[][];
    }

    export interface SqlContainerGetPropertiesResponseResource {
        /**
         * The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the container
         */
        indexingPolicy?: outputs.documentdb.IndexingPolicyResponse;
    }

}