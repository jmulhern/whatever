
export default function Track({id, url, title}) {
  const iframeStyle = {
    border: 0,
    width: '350px',
    height: '350px',
  };
  const playerUrl = "https://bandcamp.com/EmbeddedPlayer/track="+id+"/size=large/bgcol=ffffff/linkcol=0687f5/minimal=true/transparent=true/"
  return (
    <iframe style={iframeStyle} src={playerUrl} seamless>
      <a href={url}>{title}</a>
    </iframe>
  )
}