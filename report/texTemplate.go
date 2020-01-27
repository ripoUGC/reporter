/*
   Copyright 2016 Vastech SA (PTY) LTD

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package report

const defaultTemplate = `
%use square brackets as golang text templating delimiters
\documentclass{article}
\usepackage{graphicx, xcolor}
\usepackage[margin=0.4in,footskip=-0.5cm]{geometry}
\graphicspath{ {images/} }
\definecolor{antiquewhite}{rgb}{0.99, 0.96, 0.92}

\begin{document}
\pagecolor{antiquewhite}
\title{[[.Title]] [[if .VariableValues]] \\ \large [[.VariableValues]] [[end]] [[if .Description]] \\ \small [[.Description]] [[end]] \vspace{-1.0cm}}
\date{[[.FromFormatted]] to [[.ToFormatted]]}
\maketitle
\begin{center}%
[[range .Panels]][[if .IsPartialWidth]][[if .IsSingleStat]][[else]]\begin{minipage}{[[.Width]]\textwidth}%
\includegraphics[width=\textwidth]{image[[.Id]]}%
\end{minipage}\allowbreak\hspace{6px}\vspace{6px}[[end]]%
[[else]]\begin{minipage}{[[.Width]]\textwidth}%
\includegraphics[width=\textwidth]{image[[.Id]]}%
\end{minipage}\vspace{6px}%
[[end]][[end]]
\end{center}
\end{document}
`
