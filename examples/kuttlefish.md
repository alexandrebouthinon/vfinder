<h2>VFinder report</h2>

<table>
  <tr align="right">
    <td>HTML files :newspaper:</td>
    <td>{{.nbFilesFound}}</td>
  </tr>
  <tr align="right">
    <td>URLs scanned :mag:</td>
    <td>{{.nbScannedURLs}}</td>
  </tr>
  <tr align="right">
    <td>URLs errored :heavy_multiplication_x:</td>
    <td>{{.nbErroredURLs}}</td>
  </tr>
</table>

{{if (gt .nbErroredURLs "0") }}
  :warning: <strong>There are invalid links in analyzed files. Please
  check VFinder execution output for more information.</strong>

{{end}}
