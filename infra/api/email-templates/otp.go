package templates

var LoginEmailTemplateOTP = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta http-equiv="x-ua-compatible" content="ie=edge" />
    <title>Login - MFG</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
      /**
    * Better box sizing
    */
      *,
      *::before,
      *::after {
        box-sizing: border-box;
      }
      /**
    * Avoid browser level font resizing.
    * 1. Windows Mobile
    * 2. iOS / OSX
    */
      body,
      table,
      td,
      a {
        -ms-text-size-adjust: 100%; /* 1 */
        -webkit-text-size-adjust: 100%; /* 2 */
      }
      /**
    * Remove extra space added to tables and cells in Outlook.
    */
      table,
      td {
        mso-table-rspace: 0pt;
        mso-table-lspace: 0pt;
      }
      /**
    * Better fluid images in Internet Explorer.
    */
      img {
        -ms-interpolation-mode: bicubic;
      }
      /**
    * Remove blue links for iOS devices.
    */
      a[x-apple-data-detectors] {
        font-family: inherit !important;
        font-size: inherit !important;
        font-weight: inherit !important;
        line-height: inherit !important;
        color: inherit !important;
        text-decoration: none !important;
      }
      body {
        width: 100% !important;
        height: 100% !important;
        padding: 0 1rem !important;
        margin: 0 !important;
      }
      /**
    * Collapse table borders to avoid space between cells.
    */
      table {
        border-collapse: collapse !important;
      }
      a {
        color: #1f80da;
      }
      img {
        height: auto;
        line-height: 100%;
        text-decoration: none;
        border: 0;
        outline: none;
      }
    </style>
  </head>
  <body style="background-color: #e9ecef">
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tr>
        <td align="center" bgcolor="#e9ecef">
          <table
            border="0"
            cellpadding="0"
            cellspacing="0"
            width="100%"
            style="max-width: 37.5rem"
          >
            <tr>
              <td align="center" valign="top" style="padding: 2.5rem 1.5rem">
                <a
                  href="{{.RLink}}"
                  target="_blank"
                  style="display: inline-block"
                >
                  <img
                    src="./logo.png"
                    alt="Logo"
                    border="0"
                    width="112"
                    style="display: block"
                  />
                </a>
              </td>
            </tr>
          </table>
        </td>
      </tr>
      <tr>
        <td align="center" bgcolor="#e9ecef">
          <table
            border="0"
            cellpadding="0"
            cellspacing="0"
            width="100%"
            style="max-width: 37.5rem"
          >
            <tr>
              <td
                align="left"
                bgcolor="#ffffff"
                style="
                  padding: 2rem 1.5rem 0;
                  font-family: system-ui, -apple-system, BlinkMacSystemFont,
                    'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans',
                    'Helvetica Neue', sans-serif;
                  border-top-left-radius: 0.5rem;
                  border-top-right-radius: 0.5rem;
                "
              >
                <h1
                  style="
                    margin: 0;
                    font-size: 2rem;
                    font-weight: 700;
                    letter-spacing: -1px;
                    line-height: 2.5rem;
                  "
                >
                  Login Account
                </h1>
              </td>
            </tr>
            <tr>
              <td
                align="left"
                bgcolor="#ffffff"
                style="
                  padding: 1.5rem;
                  font-family: system-ui, -apple-system, BlinkMacSystemFont,
                    'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans',
                    'Helvetica Neue', sans-serif;
                  font-size: 1rem;
                  line-height: 1.5rem;
                "
              >
                <p style="margin: 0">
                  Enter the following one time code back into
                  <a href="{{.RLink}}">MFG</a>
                </p>
              </td>
            </tr>
            <tr>
              <td align="left" bgcolor="#ffffff">
                <table border="0" cellpadding="0" cellspacing="0" width="100%">
                  <tr>
                    <td
                      align="center"
                      bgcolor="#ffffff"
                      style="padding: 0.5rem 1.5rem 1.5rem 1.5rem"
                    >
                      <table border="0" cellpadding="0" cellspacing="0">
                        <tr
                          style="
                            padding: 2rem 1.5rem 0;
                            font-family: system-ui, -apple-system,
                              BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
                              Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue',
                              sans-serif;
                            border-top-left-radius: 0.5rem;
                            border-top-right-radius: 0.5rem;
                          "
                        >
                          <td align="center" style="border-radius: 0.3rem">
                            <p
                              style="
                                margin: 0;
                                font-size: 2rem;
                                font-weight: 700;
                                letter-spacing: 1.5rem;
                                line-height: 2.5rem;
                              "
                            >
                              {{.Otp}}
                            </p>
                          </td>
                        </tr>
                      </table>
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
            <tr>
              <td
                align="left"
                bgcolor="#ffffff"
                style="
                  padding: 0 1.5rem 2rem 1.5rem;
                  font-family: system-ui, -apple-system, BlinkMacSystemFont,
                    'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans',
                    'Helvetica Neue', sans-serif;
                  font-size: 1rem;
                  line-height: 1.5rem;
                  border-bottom-left-radius: 0.5rem;
                  border-bottom-right-radius: 0.5rem;
                "
              >
                <p style="margin: 0">
                  Regards,<br />
                  MFG Support
                </p>
              </td>
            </tr>
          </table>
        </td>
      </tr>
      <tr>
        <td align="center" bgcolor="#e9ecef" style="padding: 1.5rem">
          <table
            border="0"
            cellpadding="0"
            cellspacing="0"
            width="100%"
            style="max-width: 37.5rem"
          >
            <tr>
              <td
                align="center"
                bgcolor="#e9ecef"
                style="
                  padding: 0.8rem 1.5rem;
                  font-family: system-ui, -apple-system, BlinkMacSystemFont,
                    'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans',
                    'Helvetica Neue', sans-serif;
                  font-size: 0.8rem;
                  line-height: 1.25rem;
                  color: #666;
                "
              >
                <p style="margin: 0">
                  You received this email because we received a request for
                  logining into your account. If you didn't request this you can
                  safely delete this email.
                </p>
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</html>
`