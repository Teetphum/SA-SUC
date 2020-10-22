/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Eatinghistory
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntEatinghistoryEdges,
    EntEatinghistoryEdgesFromJSON,
    EntEatinghistoryEdgesFromJSONTyped,
    EntEatinghistoryEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEatinghistory
 */
export interface EntEatinghistory {
    /**
     * AddedTime holds the value of the "added_time" field.
     * @type {string}
     * @memberof EntEatinghistory
     */
    addedTime?: string;
    /**
     * 
     * @type {EntEatinghistoryEdges}
     * @memberof EntEatinghistory
     */
    edges?: EntEatinghistoryEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntEatinghistory
     */
    id?: number;
}

export function EntEatinghistoryFromJSON(json: any): EntEatinghistory {
    return EntEatinghistoryFromJSONTyped(json, false);
}

export function EntEatinghistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEatinghistory {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'added_time') ? undefined : json['added_time'],
        'edges': !exists(json, 'edges') ? undefined : EntEatinghistoryEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntEatinghistoryToJSON(value?: EntEatinghistory | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added_time': value.addedTime,
        'edges': EntEatinghistoryEdgesToJSON(value.edges),
        'id': value.id,
    };
}


