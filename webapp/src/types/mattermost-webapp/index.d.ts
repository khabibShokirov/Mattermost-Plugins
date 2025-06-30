// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

export interface PluginRegistry {
    registerPostTypeComponent(typeName: string, component: React.ElementType);

    // Add more if needed from https://developers.mattermost.com/extend/plugins/webapp/reference
}

// Добавляем объявление для файлов PNG, чтобы TypeScript их распознавал
declare module '*.png' {
    const value: string;
    export default value;
}

declare module '*.svg' {
    const content: string;
    export default content;
}
