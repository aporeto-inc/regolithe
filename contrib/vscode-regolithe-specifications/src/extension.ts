'use strict';
import * as vscode from 'vscode';
import * as path from 'path';

import { RegolitheDocumentFormattingEditProvider } from './formatter';
import { RegolitheGenerator } from './generator';


export function activate(ctx: vscode.ExtensionContext) {

    const outputChannel = vscode.window.createOutputChannel("regolithe")
    const regoPath = path.join(ctx.extensionPath, 'bin', 'rego')
    const formatter = new RegolitheDocumentFormattingEditProvider(regoPath, outputChannel);

    const regoGenFileName = '.regolithe-gen-cmd';
    const generator = new RegolitheGenerator(regoGenFileName, outputChannel);

    let lastFormatSuccess = false;

    // This is the correct way to do it but it never triggers...
    // ctx.subscriptions.push(
    //     vscode.languages.registerDocumentFormattingEditProvider(
    //         {language: 'yaml', pattern: "**/*.{spec,abs}"},
    //         formatter,
    //     ),
    // );

    ctx.subscriptions.push(
        vscode.workspace.onWillSaveTextDocument(
            (e: vscode.TextDocumentWillSaveEvent): void => {
                const res = formatter.format(e.document)
                if (res == null) {
                    return
                }
                e.waitUntil(res.then(
                    edits => {
                        lastFormatSuccess = true;
                        return edits
                    }, error => {
                        lastFormatSuccess = false;
                }))
            }
        ),
    );

    ctx.subscriptions.push(
        vscode.workspace.onDidSaveTextDocument(
            (doc: vscode.TextDocument) => {
                if (lastFormatSuccess) {
                    generator.generate(doc)
                }
                lastFormatSuccess = false;
            }
        ),
    );
}

// this method is called when your extension is deactivated
export function deactivate() {
}
